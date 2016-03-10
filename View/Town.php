<?php

/**
 * Created by IntelliJ IDEA.
 * User: Gene
 * Date: 3/8/2016
 * Time: 9:55 PM
 */
class Town
{
    private $output;
    private $town;

    public function setOutput(OutputInterface $outputType)
    {
        $this->output = $outputType;
    }

    public function loadOutput()
    {
        return $this->output->load();
    }

    public function __construct() {
        $this->town = new Town();
    }
}