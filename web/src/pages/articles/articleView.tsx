import ArticleContent from "../../components/articleContent";
import {Box, Typography} from "@mui/material";
import {ScrollRestoration} from "react-router-dom";
import React from "react";

export default function ArticleViewPage() {
    return (
    <Box>
        <ArticleContent />
        <ScrollRestoration />
    </Box>
    )
}