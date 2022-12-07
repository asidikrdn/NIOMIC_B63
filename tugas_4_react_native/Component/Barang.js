/* eslint-disable prettier/prettier */
/* eslint-disable no-alert */
/* eslint-disable react-native/no-inline-styles */
import React, {Component} from 'react';
import {Text, View, TouchableOpacity} from 'react-native';

class Barang extends Component {
  constructor(props) {
    super(props);
    this.state = {
      jumlah: 0,
    };
  }

  tambahBarang = () => {
    this.setState({jumlah: this.state.jumlah + 1});
  };

  kurangiBarang = () => {
    if (this.state.jumlah > 0) {
      this.setState({jumlah: this.state.jumlah - 1});
    } else {
      alert('Jumlah Tidak Boleh Kurang Dari 0');
    }
  };

  render() {
    return (
      <View>
        <Text style={{textAlign: 'right', marginEnd: 20}}>
          Jumlah: {this.state.jumlah}
        </Text>
        <Text style={{textAlign: 'center', fontSize: 22, marginVertical: 20}}>
          Silahkan Tekan Tombol Di Bawah
        </Text>
        <View
          style={{
            flexDirection: 'row',
            flexWrap: 'wrap',
            justifyContent: 'space-evenly',
            paddingVertical: 15,
          }}>
          <TouchableOpacity
            style={{
              backgroundColor: 'green',
              width: '30%',
              paddingVertical: 10,
            }}
            onPress={this.tambahBarang}>
            <Text style={{color: 'white', fontSize: 25, textAlign: 'center'}}>
              +
            </Text>
          </TouchableOpacity>
          <TouchableOpacity
            style={{backgroundColor: 'red', width: '30%'}}
            onPress={this.kurangiBarang}>
            <Text
              style={{
                color: 'white',
                fontSize: 25,
                textAlign: 'center',
                paddingVertical: 10,
              }}>
              -
            </Text>
          </TouchableOpacity>
        </View>
        <Text
          style={{
            textAlign: 'center',
            fontSize: 75,
            color: 'darkblue',
            marginVertical: 25,
          }}>
          {this.state.jumlah}
        </Text>
      </View>
    );
  }
}

export default Barang;
