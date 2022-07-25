<?php

session_start();
require "db.php";
if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/home.php');
  die();
}

if (!isset($_POST['env'])) {
  header('Location: /sansapp/php/mis_productos.php');
  die();
}

if (!isset($_SESSION['temp_id_producto'])) {
  header('Location: /sansapp/php/mis_productos.php');
  die();
}

$id_prod = $_SESSION['temp_id_producto'];
unset($_SESSION['temp_id_producto']);
$query = "SELECT * FROM producto WHERE id_producto='$id_prod'";
$res = mysqli_query($conn, $query);

$fila = mysqli_fetch_array($res);

if (is_null($fila)) {
  header('Location: /sansapp/php/mis_productos.php');
  die();
}

if (!($fila['vendedor'] == $_SESSION['id'])) {
  header('Location: /sansapp/php/mis_productos.php');
  die();
}

$nomb = $_POST['nombre'];
$precio = $_POST['precio'];
$etiqueta = $_POST['etiqueta'];
$descripcion = $_POST['descripcion'];

$query = "UPDATE producto SET nombre='$nomb', descripcion='$descripcion', etiqueta='$etiqueta', precio='$precio' WHERE id_producto='$id_prod'";
$res = mysqli_query($conn, $query);

header('Location: /sansapp/php/mis_productos.php');

if ($res) {
  $_SESSION['mensaje'] = "Producto modificado con exito.";
  header('Location: /sansapp/php/mis_productos.php');
} else {
  header('Location: /sansapp/php/mis_productos.php');
}
