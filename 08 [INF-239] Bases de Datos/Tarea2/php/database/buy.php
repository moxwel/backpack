<?php

session_start();
require "db.php";
if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/carrito.php');
  die();
}

if (!isset($_GET['id'])) {
  header('Location: /sansapp/php/carrito.php');
  die();
}

$id_comprador = $_GET['id'];
$query = "SELECT productos_carrito.*, producto.vendedor FROM productos_carrito JOIN producto ON productos_carrito.producto = producto.id_producto WHERE comprador='$id_comprador'";
$res = mysqli_query($conn, $query);

$fila = mysqli_fetch_array($res);

if (!($fila['comprador'] == $_SESSION['id'])) {
  header('Location: /sansapp/php/carrito.php');
  die();
}


while (!is_null($fila)) {
  $vendedor = $fila['vendedor'];
  $producto = $fila['producto'];
  $comprador = $fila['comprador'];
  $cantidad = $fila['cantidad'];
  $precio_total = $fila['precio_total'];

  $query2 = "INSERT INTO transaccion(vendedor, producto, comprador, cantidad, precio_total) VALUES ('$vendedor', '$producto', '$comprador', '$cantidad', '$precio_total')";
  $res2 = mysqli_query($conn, $query2);

  $fila = mysqli_fetch_array($res);
}

$query = "DELETE FROM productos_carrito WHERE comprador='$comprador'";
$res = mysqli_query($conn, $query);

header('Location: /sansapp/php/mi_historial.php');
