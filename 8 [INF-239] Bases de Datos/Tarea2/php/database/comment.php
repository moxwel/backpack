<?php

session_start();
require "db.php";
if (!isset($_SESSION['logged'])) {
  $_SESSION['mensaje'] = "Debes estar logueado para poder comentar.";
  header('Location: /sansapp/php/home.php');
  die();
}

if (!isset($_POST['env'])) {
  $_SESSION['mensaje'] = "Usa el formulario por favor.";
  header('Location: /sansapp/php/home.php');
  die();
} else {

  if (!isset($_SESSION['temp_id_producto'])) {
    $_SESSION['mensaje'] = "Usa el formulario por favor.";
    header('Location: /sansapp/php/home.php');
    die();
  }

  $id_prod = $_SESSION['temp_id_producto'];
  unset($_SESSION['temp_id_producto']);
  $id_cuenta = $_SESSION['id'];
  $titulo = $_POST['titulo'];
  $desc = $_POST['desc'];

  $query = "INSERT INTO comentario(producto,autor,titulo,descripcion) VALUES ('$id_prod','$id_cuenta','$titulo','$desc')";
  $res = mysqli_query($conn, $query);

  if ($res) {
    $_SESSION['mensaje'] = "Comentario añadido exitosamente.";
      header('Location: /sansapp/php/producto.php?id='.$id_prod);
  } else {
    $_SESSION['mensaje'] = "Ocurrio un error.";
    header('Location: /sansapp/php/producto.php?id='.$id_prod);
  }
}
