<?php
session_start();
require "./database/db.php";

if (!isset($_GET['id'])) {
  $_SESSION['mensaje'] = "No se seleccionó ningun producto.";
  header('Location: /sansapp/php/home.php');
  die();
}

if (!isset($_SESSION['logged'])) {
  unset($_SESSION['mensaje']);
  $_SESSION['mensaje'] = "Debes iniciar sesion para poder comentar.";
  header('Location: /sansapp/php/login.php');
  die();
}

$id_prod = $_GET['id'];
$query = "SELECT *  FROM producto WHERE id_producto='$id_prod'";
$res = mysqli_query($conn, $query);

$fila = mysqli_fetch_array($res);

if (is_null($fila)) {
  $_SESSION['mensaje'] = "No existe producto.";
  header('Location: /sansapp/php/home.php');
  die();
}

$_SESSION['temp_id_producto'] = $id_prod;
?>

<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title> Agregar producto </title>
<style>
Body {
  font-family: Calibri, Helvetica, sans-serif;
  background-color: #4BC3C7;
}
button {
       background-color: #E7A12F;
       width: 100%;
        color: white;
        padding: 15px;
        margin: 10px 0px;
        border: none;
        cursor: pointer;
         }
 form {
        border: 3px solid #f1f1f1;
    }
 input[type=text], input[type=number] {
        width: 100%;
        margin: 8px 0;
        padding: 12px 20px;
        display: inline-block;
        border: 2px solid lightblue;
        box-sizing: border-box;
    }
 button:hover {
        opacity: 0.7;
    }
  .opposition {
        width: auto;
        padding: 10px 18px;
        margin: 10px 5px;
    }


 .container {
        padding: 25px;
        background-color: lightblue;
    }

    textarea {
  width: 100%;
  margin: 8px 0;
  padding: 12px 20px;
  display: inline-block;
  border: 2px solid lightblue;
  box-sizing: border-box;
}

</style>
</head>
<body>
<center> <h1> <a href="/sansapp/php/home.php"> SansApp </a> </h1> </center>
    <form action="/sansapp/php/database/comment.php" method="post">
        <div class="container">
            <center> <h1> Comentando producto </h1> </center>
            <center> <h3> " <?php echo($fila['nombre']); ?> " </h3> </center>
            <?php
                if (isset($_SESSION['mensaje'])) {
                echo("<center><strong>" . $_SESSION['mensaje'] . "</strong></center><br>");
                unset($_SESSION['mensaje']);
                }
            ?>
            <label>✍ Titulo comentario : </label>
            <input type="text" maxlength="20" placeholder="Introduce el titulo de tu comentario" name="titulo" required>
            <label>📄 Descripcion : </label><br>
            <textarea name="desc" rows="20" cols="200"placeholder="Introduce tu comentario" required></textarea>
            <!--<label>Contraseña : </label>
            <input type="contraseña" placeholder="Introduce tu contraseña" name="contraseña" required>   -->
            <button type="submit" name="env">💬 Comentar</button>

            <hr style="width:75%;">
            <center><button type="button" class="opposition"> <a style="color:white" href="producto.php?id=<?php echo($_GET['id']); ?>">Atras</a></button> </center>
        </div>
    </form>
</body>
</html>
