import React, {Component} from 'react';
import {Text, View} from 'react-native';
import Barang from './Component/Barang';

const Header = () => {
  return <Text>TUGAS 2 REACT NATIVE</Text>;
};

class App extends Component {
  render() {
    return (
      <View>
        <Header />
        <Barang namabelanja="================ Total Belanja ================" />
      </View>
    );
  }
}

export default App;
