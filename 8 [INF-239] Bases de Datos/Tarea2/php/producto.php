<?php
session_start();
require "./database/db.php";

if (!isset($_GET['id'])) {
  $_SESSION['mensaje'] = "No se seleccionó ningun producto.";
  header('Location: /sansapp/php/home.php');
  die();
}

$id_prod = $_GET['id'];

if (isset($_SESSION['logged'])) {
  $id_usuario = $_SESSION['id'];

  $queryCalif = "SELECT * FROM calificacion WHERE autor='$id_usuario' AND producto='$id_prod'";
  $resCalif = mysqli_query($conn, $queryCalif);
  $filaCalif = mysqli_fetch_array($resCalif);
} else {
  $filaCalif = null;
}

$query = "SELECT producto.*, cuenta.nombre AS nombre_vendedor FROM producto JOIN cuenta ON producto.vendedor = cuenta.id_cuenta WHERE id_producto='$id_prod'";
$res = mysqli_query($conn, $query);

$fila = mysqli_fetch_array($res);

if (is_null($fila)) {
  $_SESSION['mensaje'] = "No existe producto.";
  header('Location: /sansapp/php/home.php');
  die();
}


$queryComm = "SELECT comentario.*, cuenta.nombre AS nombre_autor FROM comentario JOIN cuenta ON comentario.autor = cuenta.id_cuenta WHERE producto='$id_prod'";
$resComm = mysqli_query($conn, $queryComm);

$queryNCalif = "SELECT count(*) AS n_calif FROM calificacion WHERE producto='$id_prod'";
$resNCalif = mysqli_query($conn, $queryNCalif);
$filaNCalif = mysqli_fetch_array($resNCalif);

$queryNComm = "SELECT count(*) AS n_comm FROM comentario WHERE producto='$id_prod'";
$resNComm = mysqli_query($conn, $queryNComm);
$filaNComm = mysqli_fetch_array($resNComm);

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
.column2 {
  float: left;
  width: 80%;
  padding: 15px;

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


<br>
<?php
  if (isset($_SESSION['mensaje'])) {
    echo("<center><strong>" . $_SESSION['mensaje'] . "</strong></center><br>");
    unset($_SESSION['mensaje']);
  }
?>
<h2>Información Producto</h2>
<div class="row">

    <div class="column1" style="background-color:#E7A12F;">
      <h2><?php echo($fila['nombre']); ?></h2>

      <table>
        <tr>
          <th>Descripcion</th>
        </tr>
        <tr>
          <td>
            <?php echo($fila['descripcion']); ?>
          </td>
        </tr>
      </table>

      <p>• Vendedor: <a href="perfil.php?id=<?php echo($fila['vendedor']); ?>"><?php echo($fila['nombre_vendedor']); ?></a></p>
      <p>• Etiqueta: <?php echo($fila['etiqueta']); ?></p>
      <p>• Vendidos: <?php echo($fila['ventas']); ?></p>
      <h3>Calificacion: <?php echo($fila['prom_calificacion']); ?> (<?php echo($filaNCalif['n_calif']);?>)</h3>
      <h1>Precio: $ <?php echo($fila['precio']); ?></h1>
      <p>Cantidad disponible: <?php if($fila['unidades']==0){echo("Sin stock");}else{echo($fila['unidades']);} ?></p>
      <div class="topnav">

        <form action="/sansapp/php/database/addCart.php" method="POST">
          <c><label></label></c>
          <d><input hidden type="number" placeholder="Precio" name="precio" value="<?php echo($fila['precio']); ?>" required></d>
          <d><input hidden type="number" placeholder="Producto" name="id_producto" value="<?php echo($id_prod); ?>" required></d>
          <d><input style="width:100px" type="number" placeholder="Cantidad" name="cantidad" min="1" max="<?php echo($fila['unidades']); ?>" required></d>
          <a><button type="submit" name="env">🛒 Añadir al carrito</button></a>
        </form>


        <form style="float:right;" action="/sansapp/php/database/calify.php" method="POST">
        <?php
        if (!is_null($filaCalif)) {
          echo("<c style=\"font-weight:bold;\"><label>Tu calificacion: ".$filaCalif['estrellas']."</label></c>");
        }

        ?>
        <d><input hidden type="number" placeholder="Producto" name="id_producto" value="<?php echo($id_prod); ?>" required></d>
          <c>
            <select name="calif">
              <option value="5">⭐⭐⭐⭐⭐</option>
              <option value="4">⭐⭐⭐⭐</option>
              <option value="3" selected>⭐⭐⭐</option>
              <option value="2">⭐⭐</option>
              <option value="1">⭐</option>
            </select>
          </c>
          <a><button type="submit" name="env">🌟 Calificar</button></a>
          <c><label></label></c>
        </form>

      </div>
    </div>

    <div class="column2" style="background-color:lightblue;">


    <div style="background-color: #ADD8E6;" class="topnav">
      <c style="float:left; font-size: 24px; font-weight: bold;">Comentarios (<?php echo($filaNComm['n_comm']);?>)</c>
      <a style="float:right;" class="active"  href="comentar.php?id=<?php echo($fila['id_producto']); ?>">💬 Comentar</a>
    </div>


    <table>
    <tr>
      <th style="width:10%">Fecha comentario</th>
      <th style="width:10%">Autor</th>
      <th style="width:15%">Titulo</th>
      <th style="width:40%">Detalle</th>

    </tr>

    <?php

    $filaComm = mysqli_fetch_array($resComm);
    while (!is_null($filaComm)) {
    ?>

    <tr>
      <td><?php echo($filaComm['fecha_comentario']); ?></td>
      <td><a href="perfil.php?id=<?php echo($filaComm['autor']); ?>"><?php echo($filaComm['nombre_autor']); ?></a></td>
      <td><?php echo($filaComm['titulo']); ?></td>
      <td><?php echo($filaComm['descripcion']); ?></td>

    </tr>

    <?php
    $filaComm = mysqli_fetch_array($resComm);
    }
    ?>


  </table>
    </div>




</div>
</body>
</html>
