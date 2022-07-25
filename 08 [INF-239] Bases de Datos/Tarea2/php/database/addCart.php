<?php
session_start();
require "db.php";

if (!isset($_SESSION['logged'])) {
  $_SESSION['mensaje'] = "Debes iniciar sesion para poder comprar.";
  header('Location: /sansapp/php/login.php');
  die();
}

$id_comprador = $_SESSION['id'];

if (isset($_POST['env'])) {
  $id_prod = $_POST['id_producto'];
  $cantidad = $_POST['cantidad'];
  $precio = $_POST['precio'];
  $prec_tot = $cantidad * $precio;


  $query = "INSERT INTO productos_carrito(comprador, producto, cantidad, precio_total) VALUES ('$id_comprador', '$id_prod', '$cantidad', 0)";
  $res = mysqli_query($conn, $query);

  if ($res) {
    $_SESSION['mensaje'] = "Producto añadido al carrito.";
    header('Location: /sansapp/php/carrito.php');
  } else {
    $_SESSION['mensaje'] = "Primero borra el producto, y vuelvelo a ingresar.";
    header('Location: /sansapp/php/carrito.php');
  }

} else {
  header('Location: /sansapp/php/home.php');
  die();
}
