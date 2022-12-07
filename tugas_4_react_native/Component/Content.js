/* eslint-disable prettier/prettier */
/* eslint-disable radix */
import React, {useState} from 'react';
import {Input, Box, Center, Button, Flex, Text} from 'native-base';

const Content = () => {
  const [angkaPertama, setAngkaPertama] = useState();
  const [angkaKedua, setAngkaKedua] = useState();
  const [operator, setOperator] = useState('+');
  const [hasil, setHasil] = useState(0);

  const hitung = () => {
    let hasilPerhitungan;
    if (operator === '+') {
      hasilPerhitungan = parseInt(angkaPertama) + parseInt(angkaKedua);
    } else if (operator === '-') {
      hasilPerhitungan = parseInt(angkaPertama) - parseInt(angkaKedua);
    } else if (operator === '/') {
      hasilPerhitungan = parseInt(angkaPertama) / parseInt(angkaKedua);
    } else if (operator === '*') {
      hasilPerhitungan = parseInt(angkaPertama) * parseInt(angkaKedua);
    }
    setHasil(hasilPerhitungan);
  };

  return (
    <Box marginTop="5">
      <Input
        px={5}
        variant="underlined"
        placeholder="Masukkan angka"
        fontSize={20}
        keyboardType="numeric"
        onChangeText={text => setAngkaPertama(text)}
      />
      <Center
        _text={{
          color: 'darkblue',
          fontWeight: 'normal',
          fontSize: 50,
        }}
        padding="5">
        {operator}
      </Center>
      <Input
        px={5}
        variant="underlined"
        placeholder="Masukkan angka"
        fontSize={20}
        keyboardType="numeric"
        onChangeText={text => setAngkaKedua(text)}
      />
      <Box px="15" my="5">
        <Center
          _text={{
            fontSize: 50,
          }}
          borderColor="gray.400"
          borderTopWidth="3.5"
          borderBottomWidth="3"
          p="25">
          {hasil}
        </Center>
      </Box>
      <Box px="2">
        <Text fontSize="20">Operator</Text>
        <Flex direction="row" justifyContent="space-around" mb="2.5" mt="3">
          <Button
            borderRadius="50"
            bg="#4050B5"
            width={'12%'}
            onPress={() => {
              setOperator('+');
            }}>
            +
          </Button>
          <Button
            borderRadius="50"
            bg="#4050B5"
            width={'12%'}
            onPress={() => {
              setOperator('-');
            }}>
            -
          </Button>
          <Button
            borderRadius="50"
            bg="#4050B5"
            width={'12%'}
            onPress={() => {
              setOperator('/');
            }}>
            /
          </Button>
          <Button
            borderRadius="50"
            bg="#4050B5"
            width={'12%'}
            onPress={() => {
              setOperator('*');
            }}>
            *
          </Button>
          <Button
            borderLeftRadius="50"
            borderRightRadius="50"
            bg="#15AA44"
            width={'25%'}
            onPress={hitung}>
            HITUNG
          </Button>
        </Flex>
      </Box>
    </Box>
  );
};

export default Content;
