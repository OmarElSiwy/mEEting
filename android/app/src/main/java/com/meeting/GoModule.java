package com.yourappname;

import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.Promise;

import go.backend.Backend; // Import your Go package

public class MyGoModule extends ReactContextBaseJavaModule {
    public MyGoModule(ReactApplicationContext reactContext) {
        super(reactContext);
    }

    @Override
    public String getName() {
        return "MyGoModule";
    }

    @ReactMethod
    public void simpleFunction(String name, Promise promise) {
        try {
            String result = Backend.simpleFunction(name);
            promise.resolve(result);
        } catch (Exception e) {
            promise.reject("Error", e.getMessage());
        }
    }
}