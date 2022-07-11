<?php

session_start();
require "db.php";

if (!isset($_SESSION['logged'])) {
  $_SESSION['mensaje'] = "Debes iniciar sesion para poder calificar.";
  header('Location: /sansapp/php/login.php');
  die();
}

$id_autor = $_SESSION['id'];

// La comprobacion de datos solo se hara si se ha enviado un
// formulario desde la pagina correspondiente.
if (isset($_POST['env'])) {
  $id_prod = $_POST['id_producto'];
  $calif = $_POST['calif'];

  $query = "INSERT INTO calificacion(autor, producto, estrellas) VALUES('$id_autor', '$id_prod', '$calif')";
      $res = mysqli_query($conn, $query);

  if ($res) {
    $_SESSION['mensaje'] = "Producto calificado exitosamente.";
      header('Location: /sansapp/php/producto.php?id='.$id_prod);
  } else {
      $_SESSION['mensaje'] = "Calificacion modificada.";

      $query = "UPDATE calificacion SET estrellas='$calif' WHERE autor='$id_autor' AND producto='$id_prod'";
      $res = mysqli_query($conn, $query);

      header('Location: /sansapp/php/producto.php?id='.$id_prod);
  }
} else {
  $_SESSION['mensaje'] = "Usa el formulario por favor.";
  header('Location: /sansapp/php/producto.php?id='.$id_prod);
}
