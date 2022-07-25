<?php
session_start();
require "./database/db.php";

if (!isset($_GET['id'])) {
  $_SESSION['mensaje'] = "No se seleccionó ningun perfil.";
  header('Location: /sansapp/php/home.php');
  die();
}

$id_cuenta = $_GET['id'];
$query = "SELECT * FROM cuenta WHERE id_cuenta='$id_cuenta'";
$res = mysqli_query($conn, $query);

$fila = mysqli_fetch_array($res);

if (is_null($fila)) {
  $_SESSION['mensaje'] = "No existe perfil.";
  header('Location: /sansapp/php/home.php');
  die();
}

$queryProd = "SELECT * FROM producto WHERE vendedor='$id_cuenta'";
$resProd = mysqli_query($conn, $queryProd);

?>

<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title> SansApp </title>
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

.topnav b.active {
  background-color: #E7A12F;
  color: black;

}
.topnav b.active:hover {
  background-color: white;
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
.column1 {
  float: left;
  width: 60%;
  padding: 15px;

}

.row:after {
  content: "";
  display: table;
  clear: both;
}

.column2 {
  float: left;
  width: 80%;
  padding: 15px;

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

<h2>Perfil</h2>
<div class="row">
<div class="column1" style="background-color:#E7A12F;">
      <h2> <?php echo($fila['nombre']); ?> </h2>
      <p>• Correo: <?php echo($fila['correo']); ?> </p>
      <p>• Telefono: <?php if ($fila['telefono'] == -1) { echo("No establecido."); } else {echo($fila['telefono']);} ?> </p>
      <p>• Numero compras: <?php echo($fila['num_compras']); ?> </p>
      <p>• Numero ventas: <?php echo($fila['num_ventas']); ?> </p>
</div>

    <div class="column2" style="background-color:lightblue;">
      <h2>Productos en venta</h2>
      <table>
        <tr>
          <th style="width:40%">Nombre producto</th>
          <th style="width:5%">Precio</th>
          <th style="width:5%">Cantidad</th>
          <th style="width:5%">Ventas</th>
          <th style="width:5%">Calificacion</th>

        </tr>


        <?php

        $filaProd = mysqli_fetch_array($resProd);
        while (!is_null($filaProd)) {
        ?>

        <tr>
          <td><a href="producto.php?id=<?php echo($filaProd['id_producto']); ?>"><?php echo($filaProd['nombre']); ?></a></td>
          <td><?php echo($filaProd['precio']); ?></td>
          <td><?php if($filaProd['unidades']==0){echo("Sin stock");}else{echo($filaProd['unidades']);} ?></td>
          <td><?php echo($filaProd['ventas']); ?></td>
          <td><?php echo($filaProd['prom_calificacion']); ?></td>

        </tr>

        <?php
        $filaProd = mysqli_fetch_array($resProd);
        }
        ?>


  </table>
    </div>

</body>
</html>
