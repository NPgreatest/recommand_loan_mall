/*
 07-09-2015
 Created by neoneggplant
 http://neoneggplants.com

 Jan 19th 2017
 edited by meremention
 added another argument to pick the picture resolution
*/
#import "capture.h"
#import <AVFoundation/AVFoundation.h>

void usage(char* cmd),takepicture(BOOL isfront, char* sessionPreset, char* filename);
int main(int argc, char **argv, char **envp) {
    if (argc == 4) {
        if (!strcmp(argv[1], "-front")) {
            if ((!strcmp(argv[2],"480")) || (!strcmp(argv[2],"720")) || (!strcmp(argv[2],"1080"))) {
                printf("activating front camera\n");
                takepicture(TRUE,argv[2],argv[3]);
            } else {
                printf("invalid option: %s\n",argv[2]);
            }
        }
        else if (!strcmp(argv[1], "-back")) {
          if ((!strcmp(argv[2],"480")) || (!strcmp(argv[2],"720")) || (!strcmp(argv[2],"1080"))) {
            printf("activating back camera\n");
            takepicture(FALSE,argv[2],argv[3]);
        }
      }
        else {
            printf("invalid option: %s\n",argv[1]);
        }
    }
    else if (argc == 3) {
            printf("Error: %s requires 3 arguments\n",argv[0]);
    }
    else if (argc == 2) {
        if ((!strcmp(argv[1], "-h")) or (!strcmp(argv[1], "-help"))) {
            usage(argv[0]);
        }
        else if (!strcmp(argv[1], "-front")) {
            printf("Error: %s requires 3 arguments\n",argv[0]);
        }
        else if (!strcmp(argv[1], "-back")) {
            printf("Error: %s requires 3 arguments\n",argv[0]);
        }
        else {
            printf("invalid option: %s\n",argv[1]);
            printf("Error: %s requires 3 arguments\n",argv[0]);
        }
    }
    else if (argc == 1) {
        usage(argv[0]);
    }
    else  {
        printf("Error: %s requires 3 arguments\n",argv[0]);
    }
	return 0;
}

void usage(char* cmd) {
    printf("Usage: %s [-front|-back] [480|720|1080] [output]\nExample: camshot -front 720 image.jpg\n",cmd);
}

void takepicture(BOOL isFront, char* sessionPreset, char* filename) {
    //temporarily disable shutter sound
    system("mv /System/Library/Audio/UISounds/photoShutter.caf /System/Library/Audio/UISounds/photoShutter.caff >/dev/null 2>&1");
    capture *cam = [[capture alloc] init];
    [cam setupCaptureSession:isFront withResolution:sessionPreset];
    [cam setfilename:[NSString stringWithFormat:@"%s" , filename]];
    [NSThread sleepForTimeInterval:0.2];
    //wait to finish
    __block BOOL done = NO;
    [cam captureWithBlock:^(UIImage *image)
     {done = YES;}];
    while (!done)
        [[NSRunLoop mainRunLoop] runUntilDate:[NSDate dateWithTimeIntervalSinceNow:0.1]];
    [cam release];
    //return noise to system :p
    system("mv /System/Library/Audio/UISounds/photoShutter.caff /System/Library/Audio/UISounds/photoShutter.caf >/dev/null 2>&1");
}
