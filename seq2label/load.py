import re
import os
import unicodedata
import pickle
import numpy as np
from config import MAX_LENGTH, save_dir
from utilities import *
# Our dataset are three dicts: user_review, business_review, user_business_EDU
# depends on the word_vocab file
PAD_token = 0
UNK_token = 1
SOS_token = 2
EOS_token = 3
SEP_token = 4
class Voc:
    def __init__(self,env):
        self.index_word = env['index_word']
        self.word_index = env['word_index']
        self.n_words = len(self.word_index)

class Data:
    def __init__(self,env, dmax=10, smax = 20):
        self.dmax = dmax  #max number of edu
        self.smax = smax   #max number of token for a edu
        self.voc = Voc(env)
        self.train = env['train']
        self.dev = env['dev']
        self.test = env['test']
        self.train_ans = env['train_ans']
        self.dev_ans = env['dev_ans']
        self.test_ans = env['test_ans']
        self.ids = env['ids']


    

def pad_to_max(seq, seq_max, pad_token=0):
    while(len(seq)<seq_max): 
        seq.append(pad_token)
    return seq[:seq_max]




def loadPrepareData(args):

    print("Start loading...")
    path = '{}/env.json'.format(save_dir)
    print(path)
    env = dictFromFileUnicode(path)
    print('loading done...')
    ##prepare review data
    data = Data(env, args.dmax, args.smax)
    # data.user_text, user_length = prep_hierarchical_data_list(data.user_text,  data.smax, data.dmax)
    # data.item_text, item_length = prep_hierarchical_data_list(data.item_text, data.smax, data.dmax)
    # length = [user_length, item_length]#, user_length2, item_length2]
    return data#, length   #voc, user_review, business_review, user_business_EDU, train_pairs, valid_pairs, test_pairs

if __name__ == '__main__':
    loadPrepareData()
    
