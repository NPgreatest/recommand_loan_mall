import torch
USE_CUDA = torch.cuda.is_available()

MAX_LENGTH = 12
MAX_OUTPUT_LENGTH = 10
teacher_forcing_ratio = 1.0
save_dir = './data'
