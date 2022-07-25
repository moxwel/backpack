<?php

session_start();
require "db.php";
session_destroy();

session_start();
$_SESSION['mensaje'] = "Sesion cerrada.";
header('Location: /sansapp/php/login.php');
