<?php

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

    public function __construct(Town $town, $generatePhysics) {

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