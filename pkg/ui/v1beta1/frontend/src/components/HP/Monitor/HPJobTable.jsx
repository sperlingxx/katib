import React from 'react';
import { connect } from 'react-redux';

import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import TableSortLabel from '@material-ui/core/TableSortLabel';

import { fetchHPJobTrialInfo } from '../../../actions/hpMonitorActions';
import { HP_MONITOR_MODULE } from '../../../constants/constants';

const styles = theme => ({
  root: {
    width: '100%',
    marginTop: theme.spacing.unit * 3,
    overflowX: 'auto',
  },
  table: {
    minWidth: 700,
  },
  trialName: {
    '&:hover': {
      cursor: 'pointer',
    },
    width: 300,
  },
  created: {
    color: theme.colors.created,
  },
  running: {
    color: theme.colors.running,
  },
  succeeded: {
    color: theme.colors.succeeded,
  },
  killed: {
    color: theme.colors.killed,
  },
  failed: {
    color: theme.colors.failed,
  },
});

class HPJobTable extends React.Component {
  fetchAndOpenDialogTrial = trialName => () => {
    this.props.fetchHPJobTrialInfo(trialName, this.props.namespace);
  };

  constructor(props) {
    super(props);
    this.state = { orderByIdx: 1, order: 'asc' };
  }

  onChangeSortHeaderIndex = headerIndex => () => {
    if (headerIndex !== undefined) {
      const isAsc = this.state.orderByIdx === headerIndex && this.state.order === 'asc';
      this.setState(isAsc ? { order: 'desc' } : { order: 'asc' });
      this.setState({ orderByIdx: headerIndex });
    }
  };

  descendingComparator = (a, b) => {
    if (b[this.state.orderByIdx] < a[this.state.orderByIdx]) {
      return -1;
    }
    if (b[this.state.orderByIdx] > a[this.state.orderByIdx]) {
      return 1;
    }
    return 0;
  };

  getComparator = () => {
    return this.state.order === 'desc'
      ? (a, b) => this.descendingComparator(a, b)
      : (a, b) => -this.descendingComparator(a, b);
  };

  stableSort = (data, comparator) => {
    const stabilizedData = data.map((el, index) => [el, index]);
    stabilizedData.sort((a, b) => {
      const order = comparator(a[0], b[0]);
      if (order !== 0) return order;
      return a[1] - b[1];
    });
    return stabilizedData.map(el => el[0]);
  };

  render() {
    const { classes } = this.props;

    return (
      <Paper className={classes.root}>
        {this.props.jobData.length > 1 && (
          <Table className={classes.table}>
            <TableHead>
              <TableRow>
                {this.props.headers.map((header, idx) => (
                  <TableCell
                    sortDirection={this.state.orderByIdx === idx ? this.state.order : false}
                    key={idx}
                  >
                    <TableSortLabel
                      active={this.state.orderByIdx === idx}
                      direction={this.state.orderByIdx === idx ? this.state.order : 'asc'}
                      onClick={this.onChangeSortHeaderIndex(idx)}
                    >
                      {header}
                    </TableSortLabel>
                  </TableCell>
                ))}
              </TableRow>
            </TableHead>
            <TableBody>
              {this.stableSort(this.props.data, this.getComparator()).map((row, idx) => (
                <TableRow key={idx}>
                  {row.map((element, index) => {
                    if (index === 0 && row[1] === 'Succeeded') {
                      return (
                        <TableCell
                          className={classes.trialName}
                          component="th"
                          scope="row"
                          onClick={this.fetchAndOpenDialogTrial(element)}
                          key={index}
                        >
                          {element}
                        </TableCell>
                      );
                    } else if (index === 1) {
                      if (element === 'Created') {
                        return (
                          <TableCell className={classes.created} key={index}>
                            {element}
                          </TableCell>
                        );
                      } else if (element === 'Running') {
                        return (
                          <TableCell className={classes.running} key={index}>
                            {element}
                          </TableCell>
                        );
                      } else if (element === 'Succeeded') {
                        return (
                          <TableCell className={classes.succeeded} key={index}>
                            {element}
                          </TableCell>
                        );
                      } else if (element === 'Killed') {
                        return (
                          <TableCell className={classes.killed} key={index}>
                            {element}
                          </TableCell>
                        );
                      } else if (element === 'Failed') {
                        return (
                          <TableCell className={classes.failed} key={index}>
                            {element}
                          </TableCell>
                        );
                      }
                    }
                    return <TableCell key={index}>{element}</TableCell>;
                  })}
                </TableRow>
              ))}
            </TableBody>
          </Table>
        )}
      </Paper>
    );
  }
}

const mapStateToProps = state => {
  return {
    jobData: state[HP_MONITOR_MODULE].jobData,
    headers: state[HP_MONITOR_MODULE].jobData[0],
    data: state[HP_MONITOR_MODULE].jobData.slice(1),
  };
};

export default connect(mapStateToProps, { fetchHPJobTrialInfo })(withStyles(styles)(HPJobTable));
