<?php

session_start();
require "db.php";
if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/mis_productos.php');
  die();
}

if (!isset($_GET['id'])) {
  header('Location: /sansapp/php/mis_productos.php');
  die();
}

$id_prod = $_GET['id'];
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

$query = "DELETE FROM producto WHERE id_producto='$id_prod'";
$res = mysqli_query($conn, $query);

header('Location: /sansapp/php/mis_productos.php');
