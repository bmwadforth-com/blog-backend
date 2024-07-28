import {Box, Paper} from "@mui/material";
import React, {useEffect, useState} from "react";
import axios from "axios";
import {CartesianGrid, Legend, Line, LineChart, XAxis, YAxis, Tooltip} from "recharts";

export default function AboutPage() {

    return (
        <Box>
            <Paper sx={{p: 4}}>
                <h1>About bmwadforth.com</h1>
            </Paper>
        </Box>
    )
}