import * as mobx from 'mobx';
import { observer } from 'mobx-react';
import * as React from 'react';
import { GetUser } from './GetUser';
import { Search } from './Search';

@observer
export class App extends React.Component {
  @mobx.observable error: string | null = null;
  @mobx.observable userName: string | null = null;
  @mobx.observable value?: string;

  @mobx.action
  handleSubmit = (userName: string) => {
    this.userName = userName;
    this.error = null;
  };

  @mobx.action
  handleError = (error: string) => {
    this.error = error;
  };

  render() {
    if (this.userName && !this.error) {
      return <GetUser userName={this.userName} onError={this.handleError} />;
    } else {
      return <Search onSubmit={this.handleSubmit} showError={this.error} />;
    }
  }
}
