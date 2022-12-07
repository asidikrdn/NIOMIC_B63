import React, {Component} from 'react';
import {NativeBaseProvider} from 'native-base';

import Header from './Component/Header';
import Content from './Component/Content';

class App extends Component {
  render() {
    return (
      <NativeBaseProvider>
        <Header />
        <Content />
      </NativeBaseProvider>
    );
  }
}

export default App;
