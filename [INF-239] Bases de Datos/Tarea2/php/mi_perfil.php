<?php
session_start();
require "./database/db.php";

if (!isset($_SESSION['logged'])) {
  header('Location: /sansapp/php/home.php');
} else {
  $email = $_SESSION['email'];
  $query = "SELECT id_cuenta, num_compras, num_ventas FROM cuenta WHERE correo='$email'";
  $res = mysqli_query($conn, $query);
  $fila = mysqli_fetch_array($res);

  $n_compras = $fila['num_compras'];
  $n_ventas  = $fila['num_ventas'];
}
?>

<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title> Mi perfil </title>
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

</style>
<style>
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
<div style="background-color: #4BC3C7;" class="topnav">
      <c style="float:left; font-size: 24px; font-weight: bold;">Mi Perfil</c>
      <a style="float:right;" class="active" href="./database/quit.php">🚪 Cerrar Sesion</a>
    </div>
<div class="row">
    </div>
    <div class="column1" style="background-color:#E7A12F;">
      <h2> <?php echo($ses_nombre); ?> </h2>
      <p>• ROL USM: <?php echo($ses_rol); ?> </p>
      <p>• Correo: <?php echo($email); ?> </p>
      <p>• Telefono: <?php if ($_SESSION['telefono'] == -1) { echo("No establecido."); } else {echo($_SESSION['telefono']);} ?> </p>
      <p>• Numero compras: <?php echo($n_compras); ?> </p>
      <p>• Numero ventas: <?php echo($n_ventas); ?> </p>
      <div class="topnav">
        <a href="perfil.php?id=<?php echo($fila['id_cuenta']); ?>">👤 Ver Perfil Publico</a>
        <a href="modificar.php">✍ Modificar perfil</a>
        <a href="mi_historial.php">📜 Historial</a>
        <a href="mis_productos.php">📦 Administrar productos</a>
        <a style="background:red;" href="borrar_cuenta.php?id=<?php echo($fila['id_cuenta']); ?>">❌ Eliminar cuenta</a>

      </div>
    </div>




</body>
</html>
