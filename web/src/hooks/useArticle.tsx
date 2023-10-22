import {articleStateSelector, IArticleState} from "../store/articles/articlesState";
import {useRecoilValueLoadable} from "recoil";
import {useEffect} from "react";

export default function useArticle(articleSlug: string): [boolean, IArticleState] {
    const article = useRecoilValueLoadable(articleStateSelector(articleSlug));

    useEffect(() => {
        if (article.state === "hasError") {
            throw new Error(article.contents.error);
        }
    }, [article]);

    return [article.state === "loading", article.contents];
}