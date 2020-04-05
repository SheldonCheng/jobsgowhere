import * as React from "react";
import styled from "styled-components";

import Main from "../../components/Main";
import Post, { Container as StyledPost } from "./components/Post";
import PostDetail from "./components/PostDetail";

const ListContainer = styled.div`
  flex: 1;
  display: flex;
  flex-direction: column;

  ${StyledPost} {
    margin-bottom: 1rem;
  }
`;

const DetailsContainer = styled.div`
  flex: 1;
`;

type PostsScreenProps = {};

function PostsScreen(props: PostsScreenProps) {
  return (
    <Main>
      <Main.Col>
        <ListContainer>
          <Post active={true} />
          <Post active={false} />
          <Post active={false} />
          <Post active={false} />
        </ListContainer>
      </Main.Col>
      <Main.Col>
        <DetailsContainer>
          <PostDetail>Details goes here…</PostDetail>
        </DetailsContainer>
      </Main.Col>
    </Main>
  );
}

export default PostsScreen;
