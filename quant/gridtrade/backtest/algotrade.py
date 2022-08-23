"""
@desc: 使用pyalgotrade进行数据回测 https://github.com/gbeced/pyalgotrade
"""

from pyalgotrade import strategy
from pyalgotrade.barfeed import quandlfeed
from pyalgotrade_tushare import tools, barfeed


class MyStrategy(strategy.BacktestingStrategy):
    def __init__(self, feed, instrument):
        super(MyStrategy, self).__init__(feed)
        self.__instrument = instrument

    def onBars(self, bars):
        bar = bars[self.__instrument]
        self.info(bar.getClose())


# Load the bar feed from the CSV file
feed = quandlfeed.Feed()
feed.addBarsFromCSV("orcl", "WIKI-ORCL-2000-quandl.csv")

# Evaluate the strategy with the feed's bars.
myStrategy = MyStrategy(feed, "orcl")
myStrategy.run()