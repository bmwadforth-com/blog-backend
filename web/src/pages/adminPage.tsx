import { Box, Paper } from "@mui/material";
import {useRecoilValue} from "recoil";
import {userStateSelector} from "../store/articles/userState";
import {Navigate} from "react-router-dom";
import NewArticle from "../components/admin/newArticle";
import Gemini from "../components/admin/gemini";

export default function AdminPage() {
    //const user = useRecoilValue(userStateSelector);

    //TODO: fix this if (!user.isLoggedIn) return <Navigate to='/login' />;

    return (
        <Box>
            <Paper sx={{p: 4}}>
                <h1>Admin Dashboard</h1>
                <h2>Gemini</h2>
                <Gemini />

                <h2>New Article</h2>
                <NewArticle />

                <h2>Existing Articles</h2>
            </Paper>
        </Box>
    )
}