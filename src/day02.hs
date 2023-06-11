import           Control.Monad

testData = "A Y\nB X\nC Z"

data Choice = Rock | Paper | Scissors | ND deriving (Show, Read, Eq)
data Game = Game Choice Choice deriving (Show, Read, Eq)

strToChoice :: String -> Choice
strToChoice "A" = Rock
strToChoice "B" = Paper
strToChoice "C" = Scissors
strToChoice "X" = Rock
strToChoice "Y" = Paper
strToChoice "Z" = Scissors
strToChoice  _  = ND

strToGame :: String -> Game
strToGame str = Game a b
  where
    a = strToChoice . head $ words str
    b = strToChoice . last $ words str


loseTo :: Choice -> Choice
loseTo Rock     = Paper
loseTo Paper    = Scissors
loseTo Scissors = Rock
loseTo ND       = ND

winTo :: Choice -> Choice
winTo = loseTo . loseTo

choiceScore :: Game -> Int
choiceScore (Game _ Rock)     = 1
choiceScore (Game _ Paper)    = 2
choiceScore (Game _ Scissors) = 3
choiceScore (Game _ ND)       = 0

winScore :: Game -> Int
winScore (Game a b)
  | loseTo(a) == b = 6
  | a == b = 3
  | otherwise = 0

calcScore :: Game -> Int
calcScore game = (choiceScore game) + (winScore game)

choseStrategy :: Game -> Game
choseStrategy (Game a Rock)     = Game a (winTo(a))
choseStrategy (Game a Paper)    = Game a a
choseStrategy (Game a Scissors) = Game a (loseTo(a))
choseStrategy (Game a ND)       = Game ND ND

solve1st :: [Game] -> Int
solve1st =  sum . map calcScore

solve2nd :: [Game] -> Int
solve2nd =  sum . map (calcScore . choseStrategy)

main = do
    let testInput = map strToGame $ lines testData
    unless ((solve1st testInput) == 15) (error "test fail for part one")
    unless ((solve2nd testInput) == 12) (error "test fail for part two")

    contents <- readFile "dat/day02.txt"
    let input = map strToGame $ lines contents

    putStrLn $ show $ solve1st input
    putStrLn $ show $ solve2nd input
