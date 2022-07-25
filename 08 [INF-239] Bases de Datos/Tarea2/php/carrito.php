<?php
session_start();
require "./database/db.php";

if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/home.php');
}

$id_usuario = $_SESSION['id'];
$query = "SELECT producto.*, productos_carrito.cantidad AS cantidad, productos_carrito.precio_total FROM producto JOIN productos_carrito ON producto.id_producto=productos_carrito.producto WHERE productos_carrito.comprador='$id_usuario';";
$res = mysqli_query($conn, $query);

$queryTot = "SELECT SUM(precio_total) AS total_carrito FROM productos_carrito WHERE comprador='$id_usuario'";
$resTot = mysqli_query($conn, $queryTot);
$filaTot = mysqli_fetch_array($resTot);
?>

<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title> Mis productos </title>
<style>
body {
  margin: 0;
  font-family: Arial, Helvetica, sans-serif;
  background-color: #4BC3C7;
}

.topnav {
  overflow: hidden;
  background-color: lightblue;
}

.topnav a {
  float: left;
  color: black;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}

.topnav a:hover {
  background-color: #E7A12F;
  color: black;
}

.topnav a.active {
  background-color: #E7A12F;
  color: black;

}
.topnav a.active:hover {
  background-color: white;
  color: black;

}

.topnav b {
  float: right;
  color: black;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}
.topnav b:hover {
  background-color: #E7A12F;
  color: black;
}

.topnav c {
  float: left;
  color: black;
  text-align: match-parent ;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}
.topnav d {
  float: left;
  color: black;
  text-align: match-parent ;
  padding: 14px 0px;
  text-decoration: none;
  font-size: 17px;
}



.column {
  float: left;
  width: 32%;
  padding: 10px;

}


.row:after {
  content: "";
  display: table;
  clear: both;
}

</style>
<style>
  table {
  font-family: arial, sans-serif;
  border-collapse: collapse;
  width: 100%;
}

td, th {
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #dddddd;
}
</style>
</head>
<body>

<?php require "./extras/topnav.php" ?>

<br>
<?php
  if (isset($_SESSION['mensaje'])) {
    echo("<center><strong>" . $_SESSION['mensaje'] . "</strong></center><br>");
    unset($_SESSION['mensaje']);
  }
?>
<div class="row">
  <div class="column" style="background-color:#ADD8E6;width:100%;">

    <div style="background-color: #ADD8E6;" class="topnav">
      <c style="float:left; font-size: 24px; font-weight: bold;">Carrito</c>
      <a style="float:right;" class="active"  href="./database/buy.php?id=<?php echo($id_usuario); ?>">🛒 Comprar</a>
      <c style="float:right; font-size: 24px; font-weight: bold;" href="publicar.php">Total carrito: $ <?php echo($filaTot['total_carrito']); ?></c>
    </div>

    <table>
      <tr>
        <th style="width:40%">Nombre producto</th>
        <th style="width:5%">Unidades en el carrito</th>
        <th style="width:5%">Precio por unidad</th>
        <th style="width:5%">Precio total</th>
        <th style="width:10%">Accion</th>

      </tr>

      <?php

      $fila = mysqli_fetch_array($res);
      while (!is_null($fila)) {
      ?>

      <tr>
        <td><a href="producto.php?id=<?php echo($fila['id_producto']); ?>"><?php echo($fila['nombre']); ?></a></td>
        <td><?php echo($fila['cantidad']); ?></td>
        <td><?php echo($fila['precio']); ?></td>
        <td><?php echo($fila['precio_total']); ?></td>
        <th>
          <div>
            <a class="active" href="borrar_producto_carrito.php?id=<?php echo($fila['id_producto']); ?>">❌ Quitar</a>

          </div>
          </th>

      </tr>

      <?php
      $fila = mysqli_fetch_array($res);
      }
      ?>


    </table>
  </div>
  </div>

</body>
</html>
