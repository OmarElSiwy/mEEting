import Foundation

@objc(MyGoModule)
class MyGoModule: NSObject {
  @objc
  func simpleFunction(_ name: String, resolver resolve: @escaping RCTPromiseResolveBlock, rejecter reject: @escaping RCTPromiseRejectBlock) -> Void {
    let result = BackendSimpleFunction(name)
    resolve(result)
  }
}