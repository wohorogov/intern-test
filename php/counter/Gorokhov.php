<?php
session_start();
$id = session_id();

if ($id != "")
{
    $CurrentTime = time();
    $LastTime = time() - 60;

    $base = "session.txt"; // Файл, находящийся в корне каталога
    $file = file($base);

    $k = 0;
    for ($i = 0; $i < sizeof($file); $i++) {  // подсчёт активных пользователей
        $line = explode("|", $file[$i]);
        if ($line[1] > $LastTime) {
            $ResFile[$k] = $file[$i];
            $k++;
        }
    }

    for ($i = 0; $i<sizeof($ResFile); $i++) {  // поиск в файле данного пользователя по идентификатору
        $line = explode("|", $ResFile[$i]);
        if ($line[0]==$id) {
            $line[1] = trim($CurrentTime)."\n";
            $is_sid = 1;
        }
        $line = implode("|", $line); $ResFile[$i] = $line;
    }

    $f = fopen($base, "w");
    for ($i = 0; $i < sizeof($ResFile); $i++) //записываем в файл данные
    {
        fputs($f, $ResFile[$i]);
    }
    fclose($f);

    if (!$is_sid) {  // если пользователь является новым (впервые посетил сайт)
        $f = fopen($base, "a-");
        $line = $id."|".$CurrentTime."/n";
        fputs($f, $line);
        fclose($f);
    }
}
?>

<?php

    //Для того, чтобы отобразить информацию о количестве активных пользователей, в предназначенном для этого месте нужно добавть строку:
    echo "Сейчас на сайте: <b>".sizeof(file($base))."</b>";
?>
