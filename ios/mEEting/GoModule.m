#import <React/RCTBridgeModule.h>

@interface RCT_EXTERN_MODULE(MyGoModule, NSObject)

RCT_EXTERN_METHOD(exampleFunction:(NSString *)name resolver:(RCTPromiseResolveBlock)resolve rejecter:(RCTPromiseRejectBlock)reject)

@end