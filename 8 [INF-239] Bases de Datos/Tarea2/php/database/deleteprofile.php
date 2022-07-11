<?php

session_start();
require "db.php";

if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/home.php');
  die();
}

if (!isset($_GET['id'])) {
  header('Location: /sansapp/php/mi_perfil.php');
  die();
}

$id_cuenta = $_GET['id'];
$query = "SELECT * FROM cuenta WHERE id_cuenta='$id_cuenta'";
$res = mysqli_query($conn, $query);

$fila = mysqli_fetch_array($res);

if (is_null($fila)) {
  header('Location: /sansapp/php/mi_perfil.php');
  die();
}

if (!($fila['id_cuenta'] == $_SESSION['id'])) {
  header('Location: /sansapp/php/mi_perfil.php');
  die();
}

$query = "DELETE FROM cuenta WHERE id_cuenta='$id_cuenta'";
$res = mysqli_query($conn, $query);


session_destroy();
session_start();
$_SESSION['mensaje'] = "Cuenta borrada.";
header('Location: /sansapp/php/home.php');
