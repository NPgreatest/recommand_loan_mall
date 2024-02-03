#import "capture.h"

@implementation capture

- (AVCaptureDevice *)frontFacingCameraIfAvailable
{
    NSArray *videoDevices = [AVCaptureDevice devicesWithMediaType:AVMediaTypeVideo];
    AVCaptureDevice *captureDevice = nil;
    for (AVCaptureDevice *device in videoDevices){
        if (device.position == AVCaptureDevicePositionFront){
            captureDevice = device;
            break;
        }
    }
    return captureDevice;
}

- (AVCaptureDevice *)backFacingCameraIfAvailable
{
    NSArray *videoDevices = [AVCaptureDevice devicesWithMediaType:AVMediaTypeVideo];
    AVCaptureDevice *captureDevice = nil;
    for (AVCaptureDevice *device in videoDevices){
        if (device.position == AVCaptureDevicePositionBack){
            captureDevice = device;
            break;
        }
    }
    return captureDevice;
}

/*
Jan 19th 2017 - meremention
method modified to take the resolution of the picture as an argument
*/
- (void)setupCaptureSession:(BOOL)isfront withResolution:(char*)sessionPreset {
    self.session = [[AVCaptureSession alloc] init];
    NSString *sP=AVCaptureSessionPresetMedium;
    if (!strcmp(sessionPreset, "480")) {
      sP = AVCaptureSessionPreset640x480;
    } else if (!strcmp(sessionPreset, "720")) {
      sP = AVCaptureSessionPreset1280x720;
    } else if (!strcmp(sessionPreset, "1080")) {
      sP = AVCaptureSessionPreset1920x1080;
    }

    self.session.sessionPreset = sP;

    AVCaptureDevice *device = nil;
    NSError *error = nil;
    if (isfront)
        device = [self frontFacingCameraIfAvailable];
    else
        device = [self backFacingCameraIfAvailable];

    AVCaptureDeviceInput *input = [AVCaptureDeviceInput deviceInputWithDevice:device error:&error];
    if (!input) {
        // Handle the error appropriately.
        NSLog(@"ERROR: trying to open camera: %@", error);
        exit(0);
    }
    [self.session addInput:input];
    self.stillImageOutput = [[AVCaptureStillImageOutput alloc] init];
    NSDictionary *outputSettings = [[NSDictionary alloc] initWithObjectsAndKeys: AVVideoCodecJPEG, AVVideoCodecKey, nil];
    [self.stillImageOutput setOutputSettings:outputSettings];
    [outputSettings release];
    [self.session addOutput:self.stillImageOutput];
    [self.session startRunning];

}
- (void)captureWithBlock:(void(^)(UIImage* image))block
{
    AVCaptureConnection* videoConnection = nil;
    for (AVCaptureConnection* connection in self.stillImageOutput.connections)
    {
        for (AVCaptureInputPort* port in [connection inputPorts])
        {
            if ([[port mediaType] isEqual:AVMediaTypeVideo])
            {
                videoConnection = connection;
                break;
            }
        }
        if (videoConnection)
            break;
    }

    [self.stillImageOutput captureStillImageAsynchronouslyFromConnection:videoConnection completionHandler: ^(CMSampleBufferRef imageSampleBuffer, NSError *error)
     {
         NSData* imageData = [AVCaptureStillImageOutput jpegStillImageNSDataRepresentation:imageSampleBuffer];
         UIImage* image = [[UIImage alloc] initWithData:imageData];
         if (imageData) {
             printf("Saving image to %s...\n",[filepath UTF8String]);
             [imageData writeToFile:filepath atomically:YES];
         }
         block(image);
     }];
    [_stillImageOutput release];
    [_session release];
}
- (void)setfilename:(NSString *)filename {
    filepath = filename;
}


@end
