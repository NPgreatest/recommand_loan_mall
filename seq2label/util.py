import torch
import torch.nn as nn
from torch.autograd import Variable
from torch import optim
import torch.nn.functional as F
import torch.backends.cudnn as cudnn
from torch.nn.utils.rnn import pack_padded_sequence
from masked_cross_entropy import *
import itertools
import random
import math
import sys
import os
from tqdm import tqdm
from load import loadPrepareData
from load import SOS_token, EOS_token, PAD_token, UNK_token
from config import MAX_LENGTH, USE_CUDA, teacher_forcing_ratio, save_dir, MAX_OUTPUT_LENGTH
from config import MAX_LENGTH, save_dir
import pickle
import logging
logging.basicConfig(level=logging.INFO)

#############################################
# generate file name for saving parameters
#############################################
def filename(reverse, obj):
	filename = ''
	if reverse:
		filename += 'reverse_'
	filename += obj
	return filename

def split_list(lst, n, chunk_size):
    parts = []
    start_idx = 0
    for i in range(n):
        end_idx = start_idx + chunk_size
        part = lst[start_idx:end_idx]
        parts.append(part)
        start_idx += chunk_size
    return parts
def zeroPadding(l, max_target_len, fillvalue=PAD_token):
    new_l = []
    for line in l:
        if len(line) > max_target_len:
            line = line[:max_target_len]
        while len(line) < max_target_len:
            line.append(fillvalue)
        new_l.append(line)
    return new_l

def binaryMatrix(l, value=PAD_token):
    m = []
    for i in range(len(l)):
        m.append([])
        for j in range(len(l[i])):
            if l[i][j] == PAD_token:
                m[i].append(0)
            else:
                m[i].append(1) # mask = 1 if not padding
    return m

# return attribute index and input pack_padded_sequence
def inputVar(data, evaluation=False):
    comment1 = [d[0] for d in data]
    comment2 = [d[1] for d in data]
    comment1 = [split_list(x,5,MAX_LENGTH) for x in comment1]
    comment2 = [split_list(x, 5,MAX_LENGTH) for x in comment2]
    c1_len,c2_len=[],[]
    for c1 in comment1:
        c1_len.append([len(x) for x in c1])
    for c2 in comment2:
        c2_len.append([len(x) for x in c2])
    for idx in range(len(comment1)):
        comment1[idx]=zeroPadding(comment1[idx],MAX_LENGTH)
        comment2[idx]=zeroPadding(comment2[idx],MAX_LENGTH)
    # user_length = [d[0] for d in input_length]
    # business_length = [d[1] for d in input_length]
    comment1_padVar = Variable(torch.LongTensor(comment1), volatile=evaluation)
    comment2_padVar = Variable(torch.LongTensor(comment2), volatile=evaluation)

    padVar = [comment1_padVar, comment2_padVar]
    comment_len = [c1_len,c2_len]
    # length = [user_length, business_length]
    return padVar, comment_len

# convert to index, add EOS, zero padding
# return output variable, mask, max length of the sentences in batch
def outputVar(l):
    # max_target_len = max([len(indexes) for indexes in l])
    padList=zeroPadding(l, MAX_OUTPUT_LENGTH)
    # padList = zeroPadding(l, MAX_OUTPUT_LENGTH)
    mask = binaryMatrix(padList)
    mask = Variable(torch.ByteTensor(mask))
    tensor =torch.LongTensor(padList)
    padVar = Variable(tensor)
    return padVar, mask



# pair_batch is a list of (input, output) with length batch_size
# sort list of (input, output) pairs by output length, reverse input
# return input, lengths for pack_padded_sequence, output_variable, mask
def batch2TrainData(pair_batch,pair_ans, evaluation=False):
    input_batch, output_batch,ids = [], [],[]
    # input_length = []
    for i in range(0,len(pair_batch),2):
        if i+1<len(pair_batch):
            input_batch.append([pair_batch[i][1], pair_batch[i+1][1]])
            # input_length.append([user_length[str(pair_batch[i][0])], business_length[str(pair_batch[i][1])]])
            output_batch.append(pair_ans[i][1] +[4]+ pair_ans[i+1][1])
        else:
            input_batch.append([pair_batch[i][1], [0]*(5*MAX_LENGTH)])
            # input_length.append([user_length[str(pair_batch[i][0])], business_length[str(pair_batch[i][1])]])
            output_batch.append(pair_ans[i][1])
        ids.append(pair_batch[i][0])
    review_input,lenght = inputVar(input_batch,  evaluation=evaluation)
    output, mask = outputVar(output_batch) # convert sentence to ids and padding
    return review_input,lenght, output, mask,ids

