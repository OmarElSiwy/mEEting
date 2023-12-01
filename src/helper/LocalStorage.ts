import AsyncStorage from '@react-native-async-storage/async-storage';

interface DataType {
  Boomarked: string;
  DarkMode: boolean;
  Theme: string;
  Drafts: string;
  LastPage: string;
};

async function storeJsonData(key: string, jsonData: JSON) {
  try {
    const jsonString = JSON.stringify(jsonData);
    await AsyncStorage.setItem(key, jsonString);
    console.log('Data successfully saved');
  } catch (e) {
    console.error('Failed to save the data to AsyncStorage', e);
  }
}

async function getJsonData(key: string) {
  try {
    const jsonString = await AsyncStorage.getItem(key);
    if (jsonString != null) {
      const jsonData = JSON.parse(jsonString);

      // Format the data as needed
      const formattedData = formatJsonData(jsonData);

      return formattedData;
    } else {
      return null; // or an appropriate default value
    }
  } catch(e) {
    console.error('Failed to fetch the data from AsyncStorage', e);
    return null; // You can return null or an appropriate error
  }
}

function formatJsonData(jsonData: any) {
  // Assuming jsonData is an object containing the fields of DataType
  const formattedData = {
    Bookmarked: jsonData.Bookmarked || 'default bookmarked value', // Replace with default or formatting as needed
    DarkMode: jsonData.DarkMode === true, // Ensures it's always a boolean
    Theme: jsonData.Theme || 'default theme', // Replace with default theme
    Drafts: jsonData.Drafts || 'default drafts value', // Replace with default or formatting as needed
    LastPage: jsonData.LastPage || 'default last page' // Replace with default or formatting as needed
  };

  return formattedData;
}
