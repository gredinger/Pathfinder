<?php

/**
 * Created by IntelliJ IDEA.
 * User: Gene
 * Date: 3/8/2016
 * Time: 10:38 PM
 */
class Database
{

    private static $database;

    private $_host;
    private $_username;
    private $_password;
    private $_database;
    private $_type;
    private $_port;
    private $_tablePrefix;

    public static function getInstance() {
        if(null === static::$database) {
            static::$database = new self();
        }
        return static::$database;
    }

    private function __construct() {
        try {
            $this->loadConfiguration($_SERVER['DOCUMENT_ROOT'].DIRECTORY_SEPARATOR."config.ini.php");
        } catch (Exception $e) {
            die("Database configuration not valid or not found.");
        }

        try {
            $this->_connection = new mysqli($this->_host, $this->_username,
                $this->_password, $this->_database, $this->_port);
            if(mysqli_connect_error()) {
                trigger_error("Failed to connect to MySQL: " . mysql_connect_error(),
                    E_USER_ERROR);
            }
        } catch (Exception $e) {
            die("Database connection failure. Stack trace: " . $e);
        }

    }
    private function loadConfiguration(File $file) {
        $ini_array = parse_ini_file($file, true);

        $this->_username = $ini_array["database"]["username"];
        $this->_type = $ini_array["database"]["type"];
        $this->_password = $ini_array["database"]["password"];
        $this->_host = $ini_array["database"]["host"];
        $this->_port = $ini_array["database"]["port"];
        $this->_tablePrefix = $ini_array["database"]["table_prefix"];
    }

}

/*
 *
 * https://webcache.googleusercontent.com/search?q=cache:DFNRQtsptRUJ:https://gist.github.com/4534794+&cd=4&hl=en&ct=clnk&gl=us
 *
 * Jack all the codes!
 */