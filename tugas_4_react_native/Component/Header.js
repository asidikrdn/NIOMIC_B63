/* eslint-disable prettier/prettier */
/* eslint-disable react-native/no-inline-styles */
import React from 'react';
import {Text, View} from 'react-native';

const Header = () => {
  return (
    <View style={{padding: 30, backgroundColor: '#00F280'}}>
      <Text
        style={{
          textAlign: 'center',
          fontSize: 20,
          fontWeight: 'bold',
        }}>
        KALKULATOR REACT NATIVE
      </Text>
    </View>
  );
};

export default Header;
