<?php
session_start();
require "./database/db.php";

if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/home.php');
}

$id_cuenta = $_SESSION['id'];

$queryCompra = "SELECT transaccion.*, producto.nombre, cuenta.nombre AS nombre_vendedor FROM transaccion JOIN producto ON transaccion.producto = producto.id_producto JOIN cuenta ON transaccion.vendedor=cuenta.id_cuenta WHERE transaccion.comprador='$id_cuenta'";
$resCompra = mysqli_query($conn, $queryCompra);

$queryVenta = "SELECT transaccion.*, producto.nombre, cuenta.nombre AS nombre_comprador FROM transaccion JOIN producto ON transaccion.producto = producto.id_producto JOIN cuenta ON transaccion.comprador=cuenta.id_cuenta WHERE transaccion.vendedor='$id_cuenta'";
$resVenta = mysqli_query($conn, $queryVenta);

?>

<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title> Mi historial </title>
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
  padding: 10px;

}


.row:after {
  content: "";
  display: table;
  clear: both;
}

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
<h2>Historial</h2>
<div class="row">
    <div class="column" style="background-color:lightblue;width=80%;">
      <h2>Mis ventas</h2>
      <table>
      <tr>
        <th style="width:110px">Fecha venta</th>
        <th style="width:110px">Vendido a</th>
        <th style="width:200px">Producto</th>
        <th style="width:100px">Cantidad</th>
        <th style="width:100px">Total</th>

      </tr>


      <?php

      $filaVenta = mysqli_fetch_array($resVenta);
      while (!is_null($filaVenta)) {
      ?>

      <tr>
        <td><?php echo($filaVenta['fecha_transacc']); ?></td>
        <td><a href="perfil.php?id=<?php echo($filaVenta['comprador']); ?>"><?php echo($filaVenta['nombre_comprador']); ?></a></td>
        <td><a href="producto.php?id=<?php echo($filaVenta['producto']); ?>"><?php echo($filaVenta['nombre']); ?></a></td>
        <td><?php echo($filaVenta['cantidad']); ?></td>
        <td><?php echo($filaVenta['precio_total']); ?></td>

      </tr>

      <?php
      $filaVenta = mysqli_fetch_array($resVenta);
      }
      ?>


    </table>

    </div>
    <div class="column" style="background-color:#E7A12F;width=80%;">
      <h2>Mis compras</h2>
      <table>
      <tr>
        <th style="width:110px">Fecha compra</th>
        <th style="width:110px">Comprado a</th>
        <th style="width:200px">Producto</th>
        <th style="width:100px">Cantidad</th>
        <th style="width:100px">Total</th>

      </tr>

      <?php

      $filaCompra = mysqli_fetch_array($resCompra);
      while (!is_null($filaCompra)) {
      ?>

      <tr>
        <td><?php echo($filaCompra['fecha_transacc']); ?></td>
        <td><a href="perfil.php?id=<?php echo($filaCompra['vendedor']); ?>"><?php echo($filaCompra['nombre_vendedor']); ?></a></td>
        <td><a href="producto.php?id=<?php echo($filaCompra['producto']); ?>"><?php echo($filaCompra['nombre']); ?></a></td>
        <td><?php echo($filaCompra['cantidad']); ?></td>
        <td><?php echo($filaCompra['precio_total']); ?></td>

      </tr>

      <?php
      $filaCompra = mysqli_fetch_array($resCompra);
      }
      ?>


    </table>

    </div>
  </div>

</body>
</html>
