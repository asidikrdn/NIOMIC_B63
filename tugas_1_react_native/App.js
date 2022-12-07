import React, {Component} from 'react';
import {StatusBar, Text, View} from 'react-native';
import MenuUtama from './Component/MenuUtama';
import MenuGame from './Component/MenuGame';
import MenuBerita from './Component/MenuBerita';
import MenuHits from './Component/MenuHits';

const Header = () => {
  return <Text>Halaman Header</Text>;
};

const Footer = () => {
  return <Text>Footer @Niomic</Text>;
};

class App extends Component {
  render() {
    return (
      <View>
        <StatusBar backgroundColor={'green'} />
        <Header />
        <MenuUtama />
        <MenuGame />
        <MenuBerita />
        <MenuHits />
        <Footer />
      </View>
    );
  }
}

export default App;
