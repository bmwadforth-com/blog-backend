import Articles from "../../components/articles";
import {Box, Typography} from "@mui/material";
import {ScrollRestoration} from "react-router-dom";
import React from "react";

export default function ArticlesPage() {
    return(
        <Box>
            <Articles />
            <ScrollRestoration />
        </Box>
    )
}