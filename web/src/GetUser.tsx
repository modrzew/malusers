import * as React from 'react';
import { Loader } from './Loader';
import { User } from './User';

const API_URL = 'https://api.mal.modriv.net';

type Props = {
  userName: string | null;
  onError(error: string): void;
};
type State = {
  name: string | null;
  isLoading: boolean;
};

export class GetUser extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      isLoading: false,
      name: null
    };
  }

  componentDidMount() {
    this.setState({ isLoading: true });

    fetch(`${API_URL}/user/${this.props.userName}`)
      .then(resp => resp.json())
      .then(data => {
        this.setState({ name: data, isLoading: false });
      })
      .catch(error => {
        this.props.onError(error);
      });
  }

  render() {
    const { name, isLoading } = this.state;
    if (isLoading) {
      return <Loader />;
    } else if (name !== null) {
      return <User name={this.state.name} />;
    } else {
      return <Loader />;
    }
  }
}
