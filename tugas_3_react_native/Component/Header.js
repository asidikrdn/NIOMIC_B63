/* eslint-disable prettier/prettier */
/* eslint-disable react-native/no-inline-styles */
import React from 'react';
import {Text, View} from 'react-native';

const Header = props => {
  return (
    <View style={{padding: 30, backgroundColor: '#9BAB4E'}}>
      <Text
        style={{
          textAlign: 'center',
          color: 'white',
          fontSize: 20,
        }}>
        Tugas 3 React Native
      </Text>
      <Text
        style={{
          textAlign: 'center',
          fontSize: 30,
          color: '#993AA1',
          fontWeight: 'bold',
        }}>
        {props.data}
      </Text>
    </View>
  );
};

export default Header;
