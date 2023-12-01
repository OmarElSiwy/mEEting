import Foundation

@objc(MyGoModule)
class MyGoModule: NSObject {
  @objc
  func exampleFunction(_ name: String, resolver resolve: @escaping RCTPromiseResolveBlock, rejecter reject: @escaping RCTPromiseRejectBlock) {
    let result = BackendsomeFunction(name)
    resolve(result)
  }

  // React Native requires this method to be present
  @objc static func requiresMainQueueSetup() -> Bool {
    return true
  }
}