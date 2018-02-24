import * as React from 'react';
import { GetUser } from './GetUser';
import { Search } from './Search';

type State = {
  userName: string | null;
  error: string | null;
  value?: string;
};

class App extends React.Component<{}, State> {
  constructor(props: {}) {
    super(props);
    this.state = {
      error: null,
      userName: null
    };
  }

  handleChange = (event: any) => {
    this.setState({ value: event.target.value });
  };

  handleSubmit = (userName: string) => {
    this.setState({ userName, error: null });
  };

  handleError = (error: string) => {
    this.setState({ error });
  };

  render() {
    const userName = this.state.userName;
    const error = this.state.error;
    if (userName && !error) {
      return <GetUser userName={userName} onError={this.handleError} />;
    } else {
      return <Search onSubmit={this.handleSubmit} showError={error} />;
    }
  }
}

export default App;
