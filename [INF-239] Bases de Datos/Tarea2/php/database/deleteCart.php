<?php

session_start();
require "db.php";
if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/carrito.php');
  die();
}

$id_usuario = $_SESSION['id'];

if (!isset($_GET['id'])) {
  header('Location: /sansapp/php/carrito.php');
  die();
}

$id_prod = $_GET['id'];
$query = "SELECT * FROM productos_carrito WHERE comprador='$id_usuario' AND producto='$id_prod'";
$res = mysqli_query($conn, $query);

$fila = mysqli_fetch_array($res);

if (is_null($fila)) {
  header('Location: /sansapp/php/carrito.php');
  die();
}

if (!($fila['comprador'] == $_SESSION['id'])) {
  header('Location: /sansapp/php/carrito.php');
  die();
}

$query = "DELETE FROM productos_carrito WHERE comprador='$id_usuario' AND producto='$id_prod'";
$res = mysqli_query($conn, $query);

header('Location: /sansapp/php/carrito.php');
