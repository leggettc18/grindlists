import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions =  {}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type CreateListItemInput = {
  name: Scalars['String'];
  source: Scalars['String'];
  quantity?: Maybe<Scalars['Int']>;
  collected: Scalars['Boolean'];
  list_id: Scalars['ID'];
};

export type Item = {
  __typename?: 'Item';
  id: Scalars['ID'];
  name: Scalars['String'];
  source: Scalars['String'];
};

export type ItemInput = {
  name: Scalars['String'];
  source: Scalars['String'];
};

export type List = {
  __typename?: 'List';
  id: Scalars['ID'];
  name: Scalars['String'];
  user: User;
  items?: Maybe<Array<ListItem>>;
};

export type ListInput = {
  name: Scalars['String'];
};

export type ListItem = {
  __typename?: 'ListItem';
  id: Scalars['ID'];
  quantity?: Maybe<Scalars['Int']>;
  collected: Scalars['Boolean'];
  list: List;
  item: Item;
};

export type ListItemInput = {
  quantity?: Maybe<Scalars['Int']>;
  collected: Scalars['Boolean'];
  list_id: Scalars['ID'];
  item_id: Scalars['ID'];
};

export type LoginInput = {
  email: Scalars['String'];
  password: Scalars['String'];
};

export type LogoutOutput = {
  __typename?: 'LogoutOutput';
  succeeded: Scalars['Boolean'];
};

export type Mutation = {
  __typename?: 'Mutation';
  login: User;
  register: User;
  refresh: User;
  logout: LogoutOutput;
  updateUser: User;
  deleteUser: User;
  createList: List;
  updateList: List;
  deleteList: List;
  createListItem: Item;
  updateItem: Item;
  deleteItem: Item;
  setListItem: ListItem;
  updateListItem: ListItem;
  unsetListItem: ListItem;
};


export type MutationLoginArgs = {
  data: LoginInput;
};


export type MutationRegisterArgs = {
  data: UserInput;
};


export type MutationUpdateUserArgs = {
  id: Scalars['ID'];
  data: UserInput;
};


export type MutationDeleteUserArgs = {
  id: Scalars['ID'];
};


export type MutationCreateListArgs = {
  data: ListInput;
};


export type MutationUpdateListArgs = {
  id: Scalars['ID'];
  data: ListInput;
};


export type MutationDeleteListArgs = {
  id: Scalars['ID'];
};


export type MutationCreateListItemArgs = {
  listItemData: CreateListItemInput;
};


export type MutationUpdateItemArgs = {
  id: Scalars['ID'];
  data: ItemInput;
};


export type MutationDeleteItemArgs = {
  id: Scalars['ID'];
};


export type MutationSetListItemArgs = {
  data: ListItemInput;
};


export type MutationUpdateListItemArgs = {
  id: Scalars['ID'];
  data: ListItemInput;
};


export type MutationUnsetListItemArgs = {
  id: Scalars['ID'];
};

export type Query = {
  __typename?: 'Query';
  me: User;
  user?: Maybe<User>;
  users: Array<User>;
  list?: Maybe<List>;
  lists: Array<List>;
  item?: Maybe<Item>;
  items: Array<Item>;
};


export type QueryUserArgs = {
  id: Scalars['ID'];
};


export type QueryListArgs = {
  id: Scalars['ID'];
};


export type QueryItemArgs = {
  id: Scalars['ID'];
};

export type User = {
  __typename?: 'User';
  id: Scalars['ID'];
  name: Scalars['String'];
  email: Scalars['String'];
  lists?: Maybe<Array<List>>;
};

export type UserInput = {
  name: Scalars['String'];
  email: Scalars['String'];
  password: Scalars['String'];
};

export type LoginMutationVariables = Exact<{
  data: LoginInput;
}>;


export type LoginMutation = (
  { __typename?: 'Mutation' }
  & { login: (
    { __typename?: 'User' }
    & Pick<User, 'id' | 'name' | 'email'>
  ) }
);


export const LoginDocument = gql`
    mutation Login($data: LoginInput!) {
  login(data: $data) {
    id
    name
    email
  }
}
    `;
export type LoginMutationFn = Apollo.MutationFunction<LoginMutation, LoginMutationVariables>;

/**
 * __useLoginMutation__
 *
 * To run a mutation, you first call `useLoginMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useLoginMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [loginMutation, { data, loading, error }] = useLoginMutation({
 *   variables: {
 *      data: // value for 'data'
 *   },
 * });
 */
export function useLoginMutation(baseOptions?: Apollo.MutationHookOptions<LoginMutation, LoginMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<LoginMutation, LoginMutationVariables>(LoginDocument, options);
      }
export type LoginMutationHookResult = ReturnType<typeof useLoginMutation>;
export type LoginMutationResult = Apollo.MutationResult<LoginMutation>;
export type LoginMutationOptions = Apollo.BaseMutationOptions<LoginMutation, LoginMutationVariables>;