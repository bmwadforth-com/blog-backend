import React, {Suspense} from 'react';
import {
    Container,
    Paper,
    createTheme,
    ThemeProvider,
    useMediaQuery, BottomNavigation, BottomNavigationAction, Tooltip
} from '@mui/material';
import {
    createBrowserRouter, createRoutesFromElements,
    Route, RouterProvider
} from "react-router-dom";
import ArticlesPage from "./pages/articles/articles";
import ArticleViewPage from "./pages/articles/articleView";
import NotFoundPage from "./pages/notFoundPage";
import { LinkedIn, GitHub } from '@mui/icons-material';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import ErrorBoundary from './pages/errorPage';
import AdminPage from './pages/adminPage';
import {useRecoilValue} from "recoil";
import userState from "./store/articles/userState";
import LoginPage from "./pages/loginPage";
import {Spinner} from "reactstrap";
import ReactGA from 'react-ga';

const TRACKING_ID = "UA-178141115-1";
ReactGA.initialize(TRACKING_ID);

export const ApplicationRoutes = {
    INDEX: '/',
    ARTICLES: '/articles',
    ARTICLE: '/article/:articleSlug',
    PROJECTS: '/projects',
    ABOUT: '/about',
    ADMIN: '/admin',
    LOGIN: '/login'
}

function App() {
    const user = useRecoilValue(userState);
    const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)');

    const theme = React.useMemo(
        () =>
            createTheme({
                palette: {
                    mode: prefersDarkMode ? 'dark' : 'light',

                },
            }),
        [prefersDarkMode],
    );

    const router = createBrowserRouter(
        createRoutesFromElements(
            <>
                <Route path={ApplicationRoutes.INDEX} element={<ArticlesPage />} />
                <Route path={ApplicationRoutes.ARTICLES} element={<ArticlesPage />} />
                <Route path={ApplicationRoutes.ARTICLE} element={<ArticleViewPage />} />
                <Route path={ApplicationRoutes.ADMIN} element={
                    <Suspense fallback={<Spinner />}><AdminPage /></Suspense>
                } />
                <Route path={ApplicationRoutes.LOGIN} element={
                    <Suspense fallback={<Spinner />}><LoginPage /></Suspense>
                } />
                <Route path="*" element={<NotFoundPage />} />
            </>
        )
    );

    return (
        <ThemeProvider theme={theme}>
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" component="div">
                        Bmwadforth<b>dot</b>com
                    </Typography>
                </Toolbar>
            </AppBar>

            <Paper id="content" square elevation={6} style={{ padding: '50px 0' }}>
                <Container>
                    <ErrorBoundary>
                        <RouterProvider router={router} />
                    </ErrorBoundary>
                </Container>
            </Paper>
            
            <BottomNavigation
                sx={{ padding: 2 }}
                showLabels
            >
                <Tooltip title="Connect on LinkedIn">
                    <BottomNavigationAction label="LinkedIn" icon={<LinkedIn />} onClick={() => window.open('https://www.linkedin.com/in/brannon-wadforth-959b06120/')} />
                </Tooltip>
                <Tooltip title="Connect on GitHub">
                    <BottomNavigationAction label="GitHub" icon={<GitHub />} onClick={() => window.open('https://github.com/bmwadforth')} />
                </Tooltip>
            </BottomNavigation>
        </ThemeProvider>
    );
}

export default App;
