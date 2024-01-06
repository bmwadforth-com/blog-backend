import {Box} from "@mui/system";
import {
    Alert,
    AlertColor,
    Button,
    Divider,
    Grid,
    IconButton, LinearProgress, Paper,
    Snackbar,
    Stack,
    TextField,
    Typography
} from "@mui/material";
import React, {useState} from "react";
import {useRecoilState} from "recoil";
import {newArticleState} from "../../store/articles/articlesState";
import ArticleApiService from "../../util/articleApiService";
import AdminApiService from "../../util/adminApiService";
import {ApplicationRoutes} from "../../App";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import moment from "moment";
import {Code} from "@bmwadforth/armor-ui";

export default function Gemini() {
    const [query, setQuery] = useState('');
    const [loading, setLoading] = useState(false);
    const [geminiData, setGeminiData] = useState('');
    const [showAlert, setShowAlert] = useState({status: '', message: ''});
    const onChange = (val: string) => setQuery(val);

    const onSubmit = async (e: any) => {
        e.preventDefault();
        const apiService = new AdminApiService();
        try {
            setLoading(true);
            const response = await apiService.queryGemini(query);
            setLoading(false);
            setGeminiData(response);
            setShowAlert({status: 'success', message: 'Successfully queried gemini'});
        } catch (e) {
            setLoading(false);
            setGeminiData('');
            setShowAlert({status: 'error', message: 'Failed to query gemini'});
        }
    }
    const handleClose = () => setShowAlert({status: '', message: ''});

    return (

        <Box
            component="form"
            sx={{
                '& .MuiTextField-root': {width: '100%', my: 2},
            }}
            noValidate
            autoComplete="off"
        >
            <Grid container style={{padding: '1em', height: '100%'}}>
                {showAlert.status &&
                    <Snackbar open={!!showAlert.status || false} autoHideDuration={6000} onClose={handleClose}>
                        <Alert variant="filled" severity={showAlert.status as AlertColor}>
                            {showAlert.message}
                        </Alert>
                    </Snackbar>}
                <Grid item xs={12}>
                    <TextField id="query" name='query' label="Gemini Query" variant="filled" type="text"
                               onChange={(e: any) => onChange(e.target.value)}/>
                    <Button sx={{width: '100%', mt: 2}} variant="contained" type='submit' onClick={onSubmit} disabled={loading}>
                        Submit
                    </Button>
                </Grid>
                <Divider/>
                <Grid item xs={12}>
                    {loading ? <LinearProgress /> : <Code data={geminiData.toString()} showLineNumbers />}
                </Grid>
            </Grid>
        </Box>
    )
}