<?php

/**
 * Created by IntelliJ IDEA.
 * User: Gene
 * Date: 3/9/2016
 * Time: 8:54 PM
 */
namespace Pathfinder;

class ArrayTable
{
    private $tableInformation;

    public function __construct(array $tableInformation) {
        $this->tableInformation = $tableInformation;
    }

    public function generateHTML() {
        echo "<table><tbody>";
        foreach($this->tableInformation as $row) {
            echo "<th>";
            foreach($row as $col) {
                echo "<td>";
                foreach($col as $data) {
                    echo $data;
                }
                echo "</td>";
            }
            echo "</th>";
        }
        echo "</tbody></table>";
    }

}
