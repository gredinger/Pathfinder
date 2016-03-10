<?php

namespace Pathfinder;

/**
 * Created by IntelliJ IDEA.
 * User: Gene
 * Date: 3/8/2016
 * Time: 9:42 PM
 */

class Town
{
    private $town;
    private $townID;
    private $townName;
    private $population;

    function getPopulation() {
        return $this->population;
    }
    function setPopulation($population) {
        $this->population = $population;
    }
    function hasPopulation(){
        return ($this->population != null && $this->population > 0);
    }

    function generateNPCS() {

    }

    function getTownHash() {

    }

    function generateRandomTown() {
        $this->population = 50;

        $this->town = base64_encode($this->population);

        return $this->town;
    }
    public function __construct($data = null) {
        if($data === null) $this->town = generateRandomTown();


    }

    public function createArray() {
        $dataArray = array("","");

        $dataArray[0][0] = array("Population");
        $dataArray[1][0] = array($this->population);
    }

    public function loadTable()
    {

    }

}