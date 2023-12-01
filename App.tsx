import React, { useEffect, useState } from 'react';
import { SafeAreaView, Text } from 'react-native';
import { NativeModules } from 'react-native';

const { MyGoModule } = NativeModules;

function App(): JSX.Element {
  const [result, setResult] = useState('');

  useEffect(() => {
    MyGoModule.someFunction('YourName')
      .then((res: string) => {
        setResult(res);
      })
      .catch((err: any) => {
        console.error(err);
      });
  }, []);

  return (
    <SafeAreaView>
      <Text>Result from Go: {result}</Text>
    </SafeAreaView>
  );
}

export default App;
