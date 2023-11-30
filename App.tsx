import React, { useEffect, useState } from 'react';
import { SafeAreaView, Text } from 'react-native';
import { NativeModules } from 'react-native';

const { GoModule } = NativeModules;

function App(): JSX.Element {
  const [result, setResult] = useState('');

  useEffect(() => {
    // Call the Go function
    GoModule.simpleFunction('React Native')
      .then(setResult)
      .catch((error: any) => console.error('Error calling Go function:', error));
  }, []);

  return (
    <SafeAreaView>
      <Text>Result from Go: {result}</Text>
    </SafeAreaView>
  );
}

export default App;
