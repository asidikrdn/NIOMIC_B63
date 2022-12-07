/* eslint-disable prettier/prettier */
import React, {Component} from 'react';
import {Text, View, TextInput} from 'react-native';

class Barang extends Component {
  constructor(props) {
    super(props);
    this.state = {
      namaBarang: '',
      jumlahBarang: '',
    };
  }

  setNamaBarang = namaBarang => {
    this.setState({namaBarang: namaBarang});
  };

  setJumlahBarang = jumlah => {
    this.setState({jumlahBarang: jumlah});
  };

  render() {
    return (
      <View>
        {/* <Text>Masukkan Nama Barang</Text> */}
        <TextInput
          placeholder="Masukkan Nama Barang"
          onChangeText={newText => this.setNamaBarang(newText)}
        />
        {/* <Text>Masukkan Jumlah</Text> */}
        <TextInput
          placeholder="Masukkan Jumlah"
          onChangeText={newText => this.setJumlahBarang(newText)}
          keyboardType="numeric"
        />
        <Text style={{margin: 'auto', textAlign: 'center'}}>
          {this.props.namabelanja}
        </Text>
        <Text>Nama Barang : {this.state.namaBarang}</Text>
        <Text>Jumlah Barang : {this.state.jumlahBarang}</Text>
      </View>
    );
  }
}

export default Barang;
