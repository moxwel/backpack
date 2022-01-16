<?php

session_start();
require "db.php";
$id_vendedor = $_SESSION['id'];

// La comprobacion de datos solo se hara si se ha enviado un
// formulario desde la pagina correspondiente.
if (isset($_POST['env'])) {
  $nomb = $_POST['nombre'];
  $precio = $_POST['precio'];
  $etiqueta = $_POST['etiqueta'];
  $cantidad = $_POST['cantidad'];
  $descripcion = $_POST['descripcion'];

  $query = "INSERT INTO producto(vendedor, etiqueta, nombre, descripcion, precio, unidades) VALUES('$id_vendedor', '$etiqueta', '$nomb', '$descripcion', '$precio', '$cantidad')";
      $res = mysqli_query($conn, $query);

  if ($res) {
    $_SESSION['mensaje'] = "Producto publicado exitosamente.";
      header('Location: /sansapp/php/mis_productos.php');
  } else {
      $_SESSION['mensaje'] = "Ocurrio un error.";
      header('Location: /sansapp/php/publicar.php');
  }
} else {
  $_SESSION['mensaje'] = "Usa el formulario por favor.";
  header('Location: /sansapp/php/publicar.php');
}
