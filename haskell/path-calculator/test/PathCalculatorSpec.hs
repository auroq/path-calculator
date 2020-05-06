module PathCalculatorSpec where

import PathCalculator
import Data.Char
import Test.Hspec
import Test.Validity


spec :: Spec
spec = do
    describe "when calculating heathrow to london" $ do
        let heathrowToLondon = [Section 50 10 30, Section 5 90 20, Section 40 2 25, Section 10 8 0]
        let expected = [(B,10),(C,30),(A,5),(C,20),(B,2),(B,8),(C,0)]
        let actual = optimalPath heathrowToLondon
        it "should return correct path" $ do
            actual `shouldBe` expected

    describe "when calculating single step" $ do
        describe "and no pre-existing paths" $ do
            let expectedA = [(C,30),(B,10)]
            let expectedB = [(B,10)]
            let section = Section 50 10 30
            let (actualA, actualB) = roadStep ([], []) section
            it "should return correct path a" $ do
                actualA `shouldBe` expectedA
            it "should return correct path b" $ do
                actualB `shouldBe` expectedB

        describe "and a has longer pre-existing path" $ do
            let preA = [(A,100)]
            let preB = [(B,50)]
            let expectedA = [(C,30),(B,10),(B,50)]
            let expectedB = [(B,10),(B,50)]
            let section = Section 50 10 30
            let (actualA, actualB) = roadStep (preA, preB) section
            it "should return correct path a" $ do
                actualA `shouldBe` expectedA
            it "should return correct path b" $ do
                actualB `shouldBe` expectedB

            describe "and c is long" $ do
                let expectedA = [(A,50),(A,100)]
                let expectedB = [(B,100),(B,50)]
                let section = Section 50 100 30
                let (actualA, actualB) = roadStep (preA, preB) section
                it "should return correct path a" $ do
                    actualA `shouldBe` expectedA
                it "should return correct path b" $ do
                    actualB `shouldBe` expectedB

        describe "and b has longer pre-existing path" $ do
            let preA = [(A,60)]
            let preB = [(B,100)]
            let expectedA = [(A,50),(A,60)]
            let expectedB = [(B,10),(B,100)]
            let section = Section 50 10 30
            let (actualA, actualB) = roadStep (preA, preB) section
            it "should return correct path a" $ do
                actualA `shouldBe` expectedA
            it "should return correct path b" $ do
                actualB `shouldBe` expectedB

            describe "and c is long" $ do
                let expectedA = [(A,50),(A,60)]
                let expectedB = [(C,30),(A,50),(A,60)]
                let section = Section 50 100 30
                let (actualA, actualB) = roadStep (preA, preB) section
                it "should return correct path a" $ do
                    actualA `shouldBe` expectedA
                it "should return correct path b" $ do
                    actualB `shouldBe` expectedB

