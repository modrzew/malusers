import * as mobx from 'mobx';
import { observer } from 'mobx-react';
import * as React from 'react';
import { Loader } from './Loader';
import { User } from './User';

const API_URL = 'https://api.mal.modriv.net';

type Props = {
  userName: string | null;
  onError(error: string): void;
};

@observer
export class GetUser extends React.Component<Props> {
  @mobx.observable isLoading: boolean = false;
  @mobx.observable name: any = null;

  componentDidMount() {
    mobx.runInAction(() => (this.isLoading = true));
    fetch(`${API_URL}/user/${this.props.userName}`)
      .then(resp => resp.json())
      .then(data => {
        mobx.runInAction(() => {
          this.name = data;
          this.isLoading = false;
        });
      })
      .catch(error => {
        this.props.onError(error);
      });
  }

  render() {
    if (this.isLoading) {
      return <Loader />;
    } else if (this.name !== null) {
      return <User name={this.name} />;
    } else {
      return <Loader />;
    }
  }
}
