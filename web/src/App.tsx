import React, { createContext, Suspense, useContext } from 'react';
import {
    Container,
    Paper,
    createTheme,
    ThemeProvider,
    useMediaQuery,
    BottomNavigation,
    BottomNavigationAction,
    Tooltip,
    Button,
    AppBar,
    Toolbar,
    Typography,
    Box,
    useTheme,
    Hidden
} from '@mui/material';
import {
    createBrowserRouter,
    createRoutesFromElements,
    Outlet,
    Route,
    RouterProvider,
    ScrollRestoration,
    useNavigate
} from "react-router-dom";
import ArticlesPage from "./pages/articles/articles";
import ArticleViewPage from "./pages/articles/articleView";
import NotFoundPage from "./pages/notFoundPage";
import { LinkedIn, GitHub, Instagram } from '@mui/icons-material';
import ErrorBoundary from './pages/errorPage';
import AdminPage from './pages/adminPage';
import { useRecoilValue } from "recoil";
import userState from "./store/articles/userState";
import LoginPage from "./pages/loginPage";
import { Spinner } from "reactstrap";
import ReactGA from 'react-ga';
import AboutPage from "./pages/aboutPage";

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

const router = createBrowserRouter(
    createRoutesFromElements(
        <Route path={ApplicationRoutes.INDEX} element={<Layout />}>
            <Route index element={<ArticlesPage />} />
            <Route path={ApplicationRoutes.ARTICLES} element={<ArticlesPage />} />
            <Route path={ApplicationRoutes.ARTICLE} element={<ArticleViewPage />} />
            {/*<Route path={ApplicationRoutes.ABOUT} element={<AboutPage />} />*/}
            <Route path={ApplicationRoutes.ADMIN} element={
                <Suspense fallback={<Spinner />}><AdminPage /></Suspense>
            } />
            <Route path={ApplicationRoutes.LOGIN} element={
                <Suspense fallback={<Spinner />}><LoginPage /></Suspense>
            } />
            <Route path="*" element={<NotFoundPage />} />
        </Route>
    )
);

export const NavigationContext = createContext<any>(() => {
    console.warn('Navigation attempted before Router is ready.');
});

function Layout() {
    const user = useRecoilValue(userState);
    const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)');
    const navigate = useNavigate();
    const theme = useTheme();


    const appliedTheme = React.useMemo(
        () =>
            createTheme({
                palette: {
                    mode: prefersDarkMode ? 'dark' : 'light',
                },
            }),
        [prefersDarkMode],
    );

    return (
        <NavigationContext.Provider value={navigate}>
            <ThemeProvider theme={appliedTheme}>
                <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
                    <AppBar position="static">
                        <Toolbar>
                            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                                Bmwadforth<b>dot</b>com
                            </Typography>
                            <Hidden smDown>
                                <Button color="inherit" href="https://www.instagram.com/bmwadforth/" target="_blank" rel="noopener noreferrer">Instagram</Button>
                                <Button color="inherit" href="https://www.linkedin.com/in/brannon-wadforth-959b06120/" target="_blank" rel="noopener noreferrer">LinkedIn</Button>
                                <Button color="inherit" href="https://github.com/bmwadforth" target="_blank" rel="noopener noreferrer">GitHub</Button>
                            </Hidden>
                        </Toolbar>
                    </AppBar>

                    <Paper component="main" id="content" square elevation={6} style={{ padding: '50px 0' }}>
                        <Container maxWidth="lg">
                            <ErrorBoundary>
                                <Outlet />
                            </ErrorBoundary>
                        </Container>
                    </Paper>

                    <BottomNavigation
                        sx={{  py: 2, width: '100%' }}
                        showLabels
                    >
                        <Tooltip title="Connect on Instagram">
                            <BottomNavigationAction label="Instagram" icon={<Instagram />} onClick={() => window.open('https://www.instagram.com/bmwadforth/')} />
                        </Tooltip>
                        <Tooltip title="Connect on LinkedIn">
                            <BottomNavigationAction label="LinkedIn" icon={<LinkedIn />} onClick={() => window.open('https://www.linkedin.com/in/brannon-wadforth-959b06120/')} />
                        </Tooltip>
                        <Tooltip title="Connect on GitHub">
                            <BottomNavigationAction label="GitHub" icon={<GitHub />} onClick={() => window.open('https://github.com/bmwadforth')} />
                        </Tooltip>
                    </BottomNavigation>
                    <ScrollRestoration />
                </Box>
            </ThemeProvider>
        </NavigationContext.Provider>
    );
}

function App() {
    return <RouterProvider router={router} fallbackElement={<Spinner />} />;
}

export default App;